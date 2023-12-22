def Merge(input_data):
    merged_df = pd.DataFrame()

    for json_str in input_data['data']:
        # Convert JSON string to Python list
        data_list = json.loads(json_str)

        # Convert list to pandas DataFrame
        df = pd.DataFrame(data_list)

        # Merge the DataFrame into the merged result
        if merged_df.empty:
            merged_df = df
        else:
            merged_df = pd.merge(merged_df, df, on=input_data['group_by'], how='outer')

    # Convert the merged DataFrame back to JSON
    merged_json = merged_df.to_json(orient='records', date_format='iso')

    return merged_json

# if __name__ == "__main__":
#     import json
#     import pandas as pd
#     import os
#     with open(os.path.join("data", "2.profit.json"), 'r') as file:
#         content1 = file.read()
#     with open(os.path.join("data", "3.margin.json"), 'r') as file:
#         content2 = file.read()

#     merged = Merge({"data": [content1, content2], "group_by": ["time", "bin"]})

#     with open(os.path.join("data", "4.merged.json"), "w") as out:
#         out.write(merged)
