def Profit(input_data):
    # Load the data into a DataFrame
    data = json.loads(input_data['data'])
    df = pd.DataFrame(data)

    # Filter for "SALE:%" and "INVESTMENT:COGS:%"
    sale_condition = df['financial_entry_type'].str.startswith('SALE:')
    investment_condition = df['financial_entry_type'].str.startswith('INVESTMENT:COGS:')
    
    # Create a new column 'profit_type' to differentiate between sales and investment
    df['profit_type'] = df['financial_entry_type'].apply(lambda x: 'sale' if x.startswith('SALE:') else ('investment' if x.startswith('INVESTMENT:COGS:') else 'other'))

    # Group by 'group_by' columns and 'profit_type', then sum the amounts
    grouped = df.groupby(input_data['group_by'] + ['profit_type']).sum().reset_index()

    # Pivot the table to have 'profit_type' as columns
    pivot_df = grouped.pivot_table(index=input_data['group_by'], columns='profit_type', values='amount', aggfunc='sum').fillna(0)

    # Calculate the 'profit'
    pivot_df['profit'] = pivot_df.get('sale', 0) - pivot_df.get('investment', 0)

    # Reset index to turn the group_by fields back into columns
    result_df = pivot_df.reset_index()

    # Select only the required columns including 'group_by' and 'profit'
    required_columns = input_data['group_by'] + ['profit']
    result_df = result_df[required_columns]

    return result_df.to_json(orient='records', date_format='iso')

# if __name__ == "__main__":
#     import json
#     import pandas as pd
#     import os
#     with open(os.path.join("data", "1.series.json"), 'r') as file:
#         content = file.read()

#     profit = Profit({"data": content, "group_by": ["time", "bin"], "amount": "amount"})

#     with open(os.path.join("data", "2.profit.json"), "w") as out:
#         out.write(profit)
