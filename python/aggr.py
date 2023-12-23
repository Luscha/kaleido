def Aggr(input):
    data = json.loads(input["data"])

    # Create DataFrame from data
    df = pd.DataFrame(data)

    # Group by specified fields and resample to daily frequency
    grouped = df.groupby(input["group_by"])

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
        aggregated = grouped.agg({result_column: aggregate_function}).reset_index()

    return aggregated.to_json(orient='records', date_format='iso')

# if __name__ == "__main__":
#     import json
#     import pandas as pd
#     import os

#     with open(os.path.join("python", "sales.json"), 'r') as file:
#         content = file.read()

#     series = Aggr({"data": content, "time_column": "time", "group_by": ["customer"], "aggr": "count"})
#     print(series)

#     with open(os.path.join("python", "sales_out.json"), "w") as out:
#         out.write(series)
