def Enrich(params):
    # Parse the JSON strings into DataFrames
    base_df = pd.DataFrame(json.loads(params['data'][0]))
    add_df = pd.DataFrame(json.loads(params['data'][1]))

    # Identify the key column in the base DataFrame
    base_key = params['base_key_column']

    # Initialize an empty DataFrame for the merged data
    merged_df = pd.DataFrame()

    # Iterate over the key candidates and merge
    for key_candidate in params['add_key_candidates']:
        if key_candidate in add_df.columns:
            merged_df = base_df.merge(add_df, left_on=base_key, right_on=key_candidate, how='left')
            if not merged_df.empty:
                break

    # Select the specified columns to be added
    add_columns = params['add_columns']
    for col, default in zip(add_columns, params['default']):
        if col in add_df.columns:
            merged_df[col] = merged_df[col].fillna(default)
        else:
            merged_df[col] = default

    # Drop columns not in the base or specified in add_columns
    merged_df = merged_df[base_df.columns.tolist() + add_columns]
    return merged_df.to_json(orient='records', date_format='iso')

# if __name__ == "__main__":
#     import json
#     import pandas as pd
#     import os
#     with open(os.path.join("python", "sales_financial.json"), 'r') as file:
#         content = file.read()

#     series = Enrich({
#         "data": [
#             '[{"country": "NI", "profit": 1234}]', # base JSON string
#             '[{"pop_est": 6545502, "iso_a2": "NI"}]' # add JSON string
#         ],
#         "base_key_column": "country",
#         "add_key_candidates": ["iso_a2", "iso_a2_eh"],
#         "add_columns": ["pop_est"],
#         "default": [0]
#     })
#     print(series)

#     with open(os.path.join("python", "sales_financial_out.json"), "w") as out:
#         out.write(series)