def Arithm(params):
    from functools import reduce
     # Load the JSON data into a DataFrame
    df = pd.DataFrame(json.loads(params["data"]))

    # Group the DataFrame if 'group_by' is specified
    if 'group_by' in params and params['group_by']:
        df = df.groupby(params['group_by'], as_index=False)

    # Define arithmetic operations
    operations = {
        'div': lambda x, y: x / y,
        'add': lambda x, y: x + y,
        'sub': lambda x, y: x - y,
        'mul': lambda x, y: x * y
    }

    if params["op"] not in operations:
        raise ValueError("Invalid arithmetic operation specified")

    # Apply the arithmetic operation sequentially across the columns
    operation = operations[params["op"]]
    df[params["result"]] = reduce(operation, [df[col] for col in params["columns"]])

    # Replace inf/-inf with NaN, then replace NaN with the default value
    df[params["result"]].replace([float('inf'), float('-inf')], pd.NA, inplace=True)
    # Replace NaN in result column with the default value
    df[params["result"]].fillna(params["default"], inplace=True)

    return df.to_json(orient='records', date_format='iso')

# if __name__ == "__main__":
#     import json
#     import pandas as pd
#     import os
#     with open(os.path.join("data", "2.profit.json"), 'r') as file:
#         content1 = file.read()
#     with open(os.path.join("data", "3.margin.json"), 'r') as file:
#         content2 = file.read()

#     merged = Arithm({"data": content, "columns": ["amount", "customer"], op: "div"})

#     with open(os.path.join("data", "4.merged.json"), "w") as out:
#         out.write(merged)
