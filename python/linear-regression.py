def LinearRegression(params):
    from sklearn.linear_model import LinearRegression as lr
    # Load JSON data into DataFrame
    json_data = params["data"]
    df = pd.read_json(json_data)

    # Convert 'time' to datetime and then to ordinal
    time_column = params["time_column"]
    df[time_column] = pd.to_datetime(df[time_column]).apply(lambda x: x.toordinal())

    # Parameters for grouping and metrics
    group_by = params["group_by"]
    metrics = params["metrics"]

    # Initialize a list to store predictions
    predictions = []

    # Grouping the data
    for _, group_df in df.groupby(group_by):
        last_day = group_df[time_column].max()
        future_days = np.array([last_day + i for i in range(1, 100 + 1)]).reshape(-1, 1)
        future_days_df = pd.DataFrame(future_days, columns=[time_column])

        # # Predict each metric
        for metric in metrics:
            model = lr()
            model.fit(group_df[[time_column]], group_df[metric])
            
            # Use DataFrame with consistent feature name for prediction
            predictions_metric = model.predict(future_days_df)

            # Create prediction entries for each metric
            for i in range(100):
                prediction_entry = {
                    time_column: datetime.fromordinal(future_days[i][0]).isoformat() + 'Z',
                    metric: predictions_metric[i]
                }
                for col in group_by:
                    prediction_entry[col] = group_df[col].iloc[0]
                predictions.append(prediction_entry)

    return json.dumps(predictions)

# if __name__ == "__main__":
#     import json
#     import pandas as pd
#     import os
#     from datetime import datetime
#     import numpy as np
#     import sklearn

#     with open(os.path.join("data", "4.merged.json"), 'r') as file:
#         content = file.read()

#     params = {
#         "data": content,
#         "group_by": ["bin"],
#         "metrics": ["profit", "margin"],
#         "time_column": "time",
#         "days": 100
#     }
#     margin = LinearRegression(params)

#     with open(os.path.join("data", "5.linear-regression.json"), "w") as out:
#         out.write(margin)
