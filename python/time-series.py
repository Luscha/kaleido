# For Time-Based Frequencies:

# 'T' or 'min': Minute
# 'H': Hour
# 'D': Day
# 'B': Business Day
# 'W': Week (defaults to Sunday as the first day of the week)
# 'W-MON', 'W-TUE', etc.: Week, but starting on a specific weekday (e.g., Monday, Tuesday)
# 'M': Month end
# 'MS': Month start
# 'Q': Quarter end
# 'QS': Quarter start
# 'A', 'Y': Year end
# 'AS', 'YS': Year start
# For Custom Frequencies:

# '5T', '15T', etc.: Every 5, 15, etc., minutes
# '2H', '4H', etc.: Every 2, 4, etc., hours
# '2D', '3D', etc.: Every 2, 3, etc., days
# '2W', '3W', etc.: Every 2, 3, etc., weeks
# '2M', '3M', etc.: Every 2, 3, etc., months
# '2Q', '3Q', etc.: Every 2, 3, etc., quarters
# '2A', '2Y', etc.: Every 2, 3, etc., years
# For Business Frequencies:

# 'BM': Business month end
# 'BMS': Business month start
# 'BQ': Business quarter end
# 'BQS': Business quarter start
# 'BA', 'BY': Business year end
# 'BAS', 'BYS': Business year start
# For Anchored Frequencies:

# 'Q-JAN', 'Q-FEB', etc.: Quarterly frequency, anchored to specific end months
# 'A-JAN', 'A-FEB', etc.: Annual frequency, anchored to specific end months

# input: { data: data, time_column: string, group_by: [], op: string = SUM }
def TimeSeriesSimple(input):
    data = json.loads(input["data"])

    # Create DataFrame from data
    df = pd.DataFrame(data)

    # Convert the date field to datetime
    df[input["time_column"]] = pd.to_datetime(df[input["time_column"]])

    # Resample to daily frequency and then group by the specified fields
    # Grouping is done after resampling to ensure all groups in the same resample bucket are combined
    resampled = (df.set_index('time')
                   .groupby([pd.Grouper(freq='D')] + input["group_by"])
                   .agg({input["amount"]: 'sum'})
                   .reset_index())


    return resampled.to_json(orient='records', date_format='iso')

def TimeSeriesGapFill(input):
    data = json.loads(input["data"])

    # Create DataFrame from data
    df = pd.DataFrame(data)

    # Convert the date field to datetime
    df[input["time_column"]] = pd.to_datetime(df[input["time_column"]])

    # Set the time column as the index
    df.set_index(input["time_column"], inplace=True)

    freq = input.get("freq", "D")

    # Group by specified fields and resample to daily frequency
    grouped = df.groupby([pd.Grouper(freq=freq)] + input["group_by"])

    # Define the column name for storing the count or aggregation result
    result_column = "count" if input.get("aggr", None) == "count" else input["amount"]

    aggregate_function = input.get("aggr", "sum")
    # Set the aggregation function
    if aggregate_function == "count":
        # Directly count the rows in each group and place the result in a column named "count"
        aggregated = grouped.size().reset_index(name=result_column)
    else:
        if aggregate_function == "distinct":
            aggregate_function = lambda x: x.nunique()
        # Aggregate data
        aggregated = grouped.agg({input["amount"]: aggregate_function}).reset_index()

    print(aggregated.to_json(orient='records', date_format='iso'))

    # Generate a complete date range
    date_range = pd.date_range(start=aggregated[input["time_column"]].min(), 
                               end=aggregated[input["time_column"]].max(), 
                               freq=freq)
   
    complete_df = None
    if len(input.get("group_by", [])):
        # Create a MultiIndex from the product of date range and other group-by fields
        multi_index = pd.MultiIndex.from_product([date_range] + [df[g].unique() for g in input["group_by"]], 
                                                names=[input["time_column"]] + input["group_by"])
        # Reindex the aggregated DataFrame
        complete_df = aggregated.set_index([input["time_column"]] + input["group_by"]).reindex(multi_index, fill_value=0).reset_index()
    else:
        # If group_by is empty, use the simple date range index for reindexing
        complete_df = aggregated.set_index(input["time_column"]).reindex(date_range, fill_value=0).reset_index()
        complete_df.rename(columns={'index': input["time_column"]}, inplace=True)

    return complete_df.to_json(orient='records', date_format='iso')

# if __name__ == "__main__":
#     import json
#     import pandas as pd
#     import os
#     with open(os.path.join("data", "0.sample.json"), 'r') as file:
#         content = file.read()

#     series = TimeSeriesGapFill({"data": content, "time_column": "time", "group_by": ["financial_entry_type", "bin"], "amount": "amount"})
#     print(series)

#     with open(os.path.join("data", "1.series.json"), "w") as out:
#         out.write(series)
