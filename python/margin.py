def Margin(input_data):
    # Load and parse the JSON data
    data = json.loads(input_data['data'])
    df = pd.DataFrame(data)

    # Filtering based on financial_entry_type
    sales_df = df[df['financial_entry_type'].str.startswith('SALE:')]
    investment_df = df[df['financial_entry_type'].str.startswith('INVESTMENT:COGS:')]

    # Grouping and summing amounts
    sales_grouped = sales_df.groupby(input_data['group_by']).sum()
    investment_grouped = investment_df.groupby(input_data['group_by']).sum()

    # Merging the sales and investment dataframes
    merged_df = sales_grouped.merge(investment_grouped, left_index=True, right_index=True, how='outer', suffixes=('_sales', '_investment')).fillna(0)

    # Calculating margin
    # Avoid null margins by checking if sales amount is zero
    merged_df['margin'] = merged_df.apply(lambda row: (float(row['amount_sales']) - float(row['amount_investment'])) / float(row['amount_sales']) if row['amount_sales'] != 0 else 0, axis=1)

    # Reset index to include group_by fields in the output
    result_df = merged_df.reset_index()[input_data['group_by'] + ['margin']]

    return result_df.to_json(orient='records', date_format='iso')

# if __name__ == "__main__":
#     import json
#     import pandas as pd
#     import os
#     with open(os.path.join("data", "1.series.json"), 'r') as file:
#         content = file.read()

#     margin = Margin({"data": content, "group_by": ["time", "bin"], "amount": "amount"})

#     with open(os.path.join("data", "3.margin.json"), "w") as out:
#         out.write(margin)
