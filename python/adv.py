def Adv(input_data):
    # Load the data into a DataFrame
    data = json.loads(input_data['data'])
    df = pd.DataFrame(data)

    # Determine rows where 'financial_entry_type' starts with 'ADV'
    is_adv = df['financial_entry_type'].str.startswith('INVESTMENT:ADVERTISING:')

    # Add a column 'adv_amount' which contains 'amount' where 'is_adv' is True, else 0
    df['adv_amount'] = df['amount'].where(is_adv, 0)

    # Group by the specified columns and sum the 'adv_amount' for these groups
    grouped = df.groupby(input_data['group_by'])['adv_amount'].sum().reset_index()

    # Rename the 'adv_amount' column to 'adv' in the grouped DataFrame
    grouped.rename(columns={'adv_amount': 'adv'}, inplace=True)

    # Return the DataFrame
    return grouped.to_json(orient='records', date_format='iso')

# if __name__ == "__main__":
#     import json
#     import pandas as pd
#     import os
#     with open(os.path.join("data", "1.series.json"), 'r') as file:
#         content = file.read()

#     profit = Profit({"data": content, "group_by": ["time", "bin"], "amount": "amount"})

#     with open(os.path.join("data", "2.profit.json"), "w") as out:
#         out.write(profit)
