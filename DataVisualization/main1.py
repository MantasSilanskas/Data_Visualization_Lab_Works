#Lab work Nr.3

import glob
import os
import pandas as pd
import numpy
from scipy.spatial.distance import cityblock
from scipy.spatial import distance
from pandas.io.formats import style
import plotly_express as px
import plotly.graph_objects as go
import dash
import dash_core_components as dcc
import dash_html_components as html
from dash.dependencies import Input, Output
from sklearn.decomposition import PCA

def manhettan(df1, df2):
    distance = cityblock(df1[:1000], df2[:1000])

    return round(distance, 2)

def euclidean(df1, df2):
    distance = numpy.linalg.norm(df1 - df2)

    return round(distance, 2)

def cosine(df1, df2):
    result = distance.cosine(df1, df2)

    return round(result, 2)

def jaccard(df1, df2):
    result = distance.jaccard(df1, df2)

    return round(result, 2)

def filterDf(df, dfName):
    types = ('temperature', 'humidity', 'CO2')
    # use filtering + list comprehension to quickly get the series
    # # second list comprehension is not necessary, just breaking
    # # up the first one for readability
    allValues = [df[df['type'] == each]['value'] for each in types]
    allValues = [each.reset_index(drop=True) for each in allValues]
    
    # get every 3rd row, keep device-id, timestamp, timestamp-ms
    # # since getting every 3rd row, indexing is screwed up
    # # use reset_index() to fix that
    dfName = df.iloc[::3, :3].reset_index(drop=True)
    
    # add each series into newDf as a new column
    for series, name in zip(allValues, types):
     dfName[name] = series

    return dfName

# Data reader and dataframe creation
path = r'/home/mantassil/Desktop/DataVisualization/data'                     # use your path
file1 = glob.glob(os.path.join(path, "0033C822.csv"))     # advisable to use os.path.join as this makes concatenation OS independent
df_from_file1 = (pd.read_csv(f) for f in file1)
df1 = pd.concat(df_from_file1, ignore_index=True)

file2 = glob.glob(os.path.join(path, "0034C8DF.csv")) 
df_from_file2 = (pd.read_csv(f) for f in file2)
df2 = pd.concat(df_from_file2, ignore_index=True)

file3 = glob.glob(os.path.join(path, "0034C74E.csv")) 
df_from_file3 = (pd.read_csv(f) for f in file3)
df3 = pd.concat(df_from_file3, ignore_index=True)

file4 = glob.glob(os.path.join(path, "0034C91E.csv")) 
df_from_file4 = (pd.read_csv(f) for f in file4)
df4 = pd.concat(df_from_file4, ignore_index=True)

file5 = glob.glob(os.path.join(path, "0034C559.csv")) 
df_from_file5 = (pd.read_csv(f) for f in file5)
df5 = pd.concat(df_from_file5, ignore_index=True)

file6 = glob.glob(os.path.join(path, "0034C864.csv")) 
df_from_file6 = (pd.read_csv(f) for f in file6)
df6 = pd.concat(df_from_file6, ignore_index=True)

dfs = {"0033C822":df1,"0034C8DF": df2, "0034C74E" : df3, "0034C91E" : df4, "0034C559" : df5, "0034C864": df6}

cols_to_norm = ['CO2','temperature', 'humidity', 'total']
# g means dataframe is grouped and n that is normalized
gdf1 = filterDf(df1, "0033C822")
gdf1["total"] = gdf1["humidity"] + gdf1["CO2"] + gdf1["temperature"]
gdf1[cols_to_norm] = gdf1[cols_to_norm].apply(lambda x: (x - x.min()) / (x.max() - x.min()))

gdf2 = filterDf(df2, "0034C8DF")
gdf2["total"] = gdf2["humidity"] + gdf2["CO2"] + gdf2["temperature"]
gdf2[cols_to_norm] = gdf2[cols_to_norm].apply(lambda x: (x - x.min()) / (x.max() - x.min()))

gdf3 = filterDf(df3, "0034C74E")
gdf3["total"] = gdf3["humidity"] + gdf3["CO2"] + gdf3["temperature"]
gdf3[cols_to_norm] = gdf3[cols_to_norm].apply(lambda x: (x - x.min()) / (x.max() - x.min()))

gdf4 = filterDf(df4, "0034C91E")
gdf4["total"] = gdf4["humidity"] + gdf4["CO2"] + gdf4["temperature"]
gdf4[cols_to_norm] = gdf4[cols_to_norm].apply(lambda x: (x - x.min()) / (x.max() - x.min()))

gdf5 = filterDf(df5, "0034C559")
gdf5["total"] = gdf5["humidity"] + gdf5["CO2"] + gdf5["temperature"]
gdf5[cols_to_norm] = gdf5[cols_to_norm].apply(lambda x: (x - x.min()) / (x.max() - x.min()))

gdf6 = filterDf(df6, "0034C864")
gdf6["total"] = gdf6["humidity"] + gdf6["CO2"] + gdf6["temperature"]
gdf6[cols_to_norm] = gdf6[cols_to_norm].apply(lambda x: (x - x.min()) / (x.max() - x.min()))

devices = ["0033C822", "0034C8DF", "0034C74E", "0034C91E", "0034C559", "0034C864"]
data = [
    ["X", manhettan(gdf1[:1000].total, gdf2[:1000].total),manhettan(gdf1[:1000].total, gdf3[:1000].total),
    manhettan(gdf1[:1000].total, gdf4[:1000].total),manhettan(gdf1[:1000].total, gdf5[:1000].total),
    manhettan(gdf1[:1000].total, gdf6[:1000].total)
    ],
    [manhettan(gdf2[:1000].total, gdf1[:1000].total),"X",manhettan(gdf2[:1000].total, gdf3[:1000].total),
    manhettan(gdf2[:1000].total, gdf4[:1000].total),manhettan(gdf1[:1000].total, gdf5[:1000].total),
    manhettan(gdf1[:1000].total, gdf5[:1000].total)
    ],
    [manhettan(gdf3[:1000].total, gdf1[:1000].total), manhettan(gdf3[:1000].total, gdf2[:1000].total),"X",
    manhettan(gdf3[:1000].total, gdf4[:1000].total),manhettan(gdf3[:1000].total, gdf5[:1000].total),
    manhettan(gdf3[:1000].total, gdf6[:1000].total)
    ],
    [manhettan(gdf4[:1000].total, gdf1[:1000].total), manhettan(gdf4[:1000].total, gdf2[:1000].total),manhettan(gdf4[:1000].total, gdf3[:1000].total),
    "X",manhettan(gdf4[:1000].total, gdf5[:1000].total),manhettan(gdf4[:1000].total, gdf6[:1000].total)
    ],
    [manhettan(gdf5[:1000].total, gdf1[:1000].total), manhettan(gdf5[:1000].total, gdf2[:1000].total),manhettan(gdf5[:1000].total, gdf3[:1000].total),
    manhettan(gdf5[:1000].total, gdf4[:1000].total),"X",manhettan(gdf5[:1000].total, gdf6[:1000].total)
    ],
    [manhettan(gdf6[:1000].total, gdf1[:1000].total), manhettan(gdf6[:1000].total, gdf2[:1000].total),manhettan(gdf6[:1000].total, gdf3[:1000].total),
    manhettan(gdf6[:1000].total, gdf4[:1000].total),manhettan(gdf6[:1000].total, gdf5[:1000].total),"X"
    ],
]

print("Manhettan distance")
format_row = "{:>12}" * (len(devices) + 1)
print(format_row.format("", *devices))
for team, row in zip(devices, data):
 print(format_row.format(team, *row))

data = [
    ["X", euclidean(gdf1[:1000].total, gdf2[:1000].total),euclidean(gdf1[:1000].total, gdf3[:1000].total),
    euclidean(gdf1[:1000].total, gdf4[:1000].total),euclidean(gdf1[:1000].total, gdf5[:1000].total),
    euclidean(gdf1[:1000].total, gdf6[:1000].total)
    ],
    [euclidean(gdf2[:1000].total, gdf1[:1000].total),"X",euclidean(gdf2[:1000].total, gdf3[:1000].total),
    euclidean(gdf2[:1000].total, gdf4[:1000].total),euclidean(gdf1[:1000].total, gdf5[:1000].total),
    euclidean(gdf1[:1000].total, gdf5[:1000].total)
    ],
    [euclidean(gdf3[:1000].total, gdf1[:1000].total), euclidean(gdf3[:1000].total, gdf2[:1000].total),"X",
    euclidean(gdf3[:1000].total, gdf4[:1000].total),euclidean(gdf3[:1000].total, gdf5[:1000].total),
    euclidean(gdf3[:1000].total, gdf6[:1000].total)
    ],
    [euclidean(gdf4[:1000].total, gdf1[:1000].total), euclidean(gdf4[:1000].total, gdf2[:1000].total),euclidean(gdf4[:1000].total, gdf3[:1000].total),
    "X",euclidean(gdf4[:1000].total, gdf5[:1000].total),euclidean(gdf4[:1000].total, gdf6[:1000].total)
    ],
    [euclidean(gdf5[:1000].total, gdf1[:1000].total), euclidean(gdf5[:1000].total, gdf2[:1000].total),euclidean(gdf5[:1000].total, gdf3[:1000].total),
    euclidean(gdf5[:1000].total, gdf4[:1000].total),"X",euclidean(gdf5[:1000].total, gdf6[:1000].total)
    ],
    [euclidean(gdf6[:1000].total, gdf1[:1000].total), euclidean(gdf6[:1000].total, gdf2[:1000].total),euclidean(gdf6[:1000].total, gdf3[:1000].total),
    euclidean(gdf6[:1000].total, gdf4[:1000].total),euclidean(gdf6[:1000].total, gdf5[:1000].total),"X"
    ],
]

print()
print("Euclidean distance")
format_row = "{:>12}" * (len(devices) + 1)
print(format_row.format("", *devices))
for team, row in zip(devices, data):
 print(format_row.format(team, *row))

data = [
    ["X", cosine(gdf1[:1000].total, gdf2[:1000].total),cosine(gdf1[:1000].total, gdf3[:1000].total),
    cosine(gdf1[:1000].total, gdf4[:1000].total),cosine(gdf1[:1000].total, gdf5[:1000].total),
    cosine(gdf1[:1000].total, gdf6[:1000].total)
    ],
    [cosine(gdf2[:1000].total, gdf1[:1000].total),"X",cosine(gdf2[:1000].total, gdf3[:1000].total),
    cosine(gdf2[:1000].total, gdf4[:1000].total),cosine(gdf1[:1000].total, gdf5[:1000].total),
    cosine(gdf1[:1000].total, gdf5[:1000].total)
    ],
    [cosine(gdf3[:1000].total, gdf1[:1000].total), cosine(gdf3[:1000].total, gdf2[:1000].total),"X",
    cosine(gdf3[:1000].total, gdf4[:1000].total),cosine(gdf3[:1000].total, gdf5[:1000].total),
    cosine(gdf3[:1000].total, gdf6[:1000].total)
    ],
    [cosine(gdf4[:1000].total, gdf1[:1000].total), cosine(gdf4[:1000].total, gdf2[:1000].total),cosine(gdf4[:1000].total, gdf3[:1000].total),
    "X",cosine(gdf4[:1000].total, gdf5[:1000].total),cosine(gdf4[:1000].total, gdf6[:1000].total)
    ],
    [cosine(gdf5[:1000].total, gdf1[:1000].total), cosine(gdf5[:1000].total, gdf2[:1000].total),cosine(gdf5[:1000].total, gdf3[:1000].total),
    cosine(gdf5[:1000].total, gdf4[:1000].total),"X",cosine(gdf5[:1000].total, gdf6[:1000].total)
    ],
    [cosine(gdf6[:1000].total, gdf1[:1000].total), cosine(gdf6[:1000].total, gdf2[:1000].total),cosine(gdf6[:1000].total, gdf3[:1000].total),
    cosine(gdf6[:1000].total, gdf4[:1000].total),cosine(gdf6[:1000].total, gdf5[:1000].total),"X"
    ],
]


print()
print("Cosine distance")
format_row = "{:>12}" * (len(devices) + 1)
print(format_row.format("", *devices))
for team, row in zip(devices, data):
 print(format_row.format(team, *row))


 data = [
    ["X", jaccard(gdf1[:1000].total, gdf2[:1000].total),jaccard(gdf1[:1000].total, gdf3[:1000].total),
    jaccard(gdf1[:1000].total, gdf4[:1000].total),jaccard(gdf1[:1000].total, gdf5[:1000].total),
    jaccard(gdf1[:1000].total, gdf6[:1000].total)
    ],
    [jaccard(gdf2[:1000].total, gdf1[:1000].total),"X",jaccard(gdf2[:1000].total, gdf3[:1000].total),
    jaccard(gdf2[:1000].total, gdf4[:1000].total),jaccard(gdf1[:1000].total, gdf5[:1000].total),
    jaccard(gdf1[:1000].total, gdf5[:1000].total)
    ],
    [jaccard(gdf3[:1000].total, gdf1[:1000].total), jaccard(gdf3[:1000].total, gdf2[:1000].total),"X",
    jaccard(gdf3[:1000].total, gdf4[:1000].total),jaccard(gdf3[:1000].total, gdf5[:1000].total),
    jaccard(gdf3[:1000].total, gdf6[:1000].total)
    ],
    [jaccard(gdf4[:1000].total, gdf1[:1000].total), jaccard(gdf4[:1000].total, gdf2[:1000].total),jaccard(gdf4[:1000].total, gdf3[:1000].total),
    "X",jaccard(gdf4[:1000].total, gdf5[:1000].total),jaccard(gdf4[:1000].total, gdf6[:1000].total)
    ],
    [jaccard(gdf5[:1000].total, gdf1[:1000].total), jaccard(gdf5[:1000].total, gdf2[:1000].total),jaccard(gdf5[:1000].total, gdf3[:1000].total),
    jaccard(gdf5[:1000].total, gdf4[:1000].total),"X",jaccard(gdf5[:1000].total, gdf6[:1000].total)
    ],
    [jaccard(gdf6[:1000].total, gdf1[:1000].total), jaccard(gdf6[:1000].total, gdf2[:1000].total),jaccard(gdf6[:1000].total, gdf3[:1000].total),
    jaccard(gdf6[:1000].total, gdf4[:1000].total),jaccard(gdf6[:1000].total, gdf5[:1000].total),"X"
    ],
]

print()
print("Jaccard distance")
format_row = "{:>12}" * (len(devices) + 1)
print(format_row.format("", *devices))
for team, row in zip(devices, data):
 print(format_row.format(team, *row))


 app = dash.Dash(__name__)

 # App layout
app.layout = html.Div([

html.H1("Data visualization lab work number 6", style={'text-align':'center'}),

dcc.Dropdown(id="dropDown1",
    options=[
        {'label':'CO2','value':'CO2'},
        {'label':'Humidity', 'value':'humidity'},
        {'label':'Temperature', 'value':'temperature'},
],
    value='temperature'
),

html.Div(id='output_container', children=[]),
html.Br(),

dcc.Graph(id='first_chart', figure={}),

html.Br(),
dcc.Graph(id='second_chart', figure={}),
])

# Connect dash to plotly

@app.callback(
    [Output(component_id='output_container', component_property='children'),
     Output(component_id='first_chart', component_property='figure'),
     Output(component_id='second_chart', component_property='figure'),
     ],
    [Input(component_id='dropDown1', component_property='value'),
    ]
)
def update_graph(option_slctd):
    print(option_slctd)
    print(type(option_slctd))

    container = "The type of data choosen. by user was: {}".format(option_slctd)

    features = ["CO2", "humidity", "temperature"]
    X = gdf6[features]

    pca = PCA()
    compoments = pca.fit_transform(X)
    labels = {
        str(i): f"PC {i+1} ({var:.1f}%)"
        for i, var in enumerate(pca.explained_variance_ratio_ * 100)
    }

    total_var = pca.explained_variance_ratio_.sum() * 100
    # Plotly Express
    fig1 = px.scatter_matrix(
    compoments,
    labels=labels,
    dimensions=range(3),
    color=gdf6["CO2"],
    title=f'Total Explained Variance: {total_var:.2f}%',
)
    fig1.update_traces(diagonal_visible=False)

    loadings = pca.components_.T * numpy.sqrt(pca.explained_variance_)

    fig2 = px.scatter(compoments, x=0, y=1, color=gdf6["CO2"])

    for i, feature in enumerate(features):
        fig2.add_shape(
            type='line',
            x0=0, y0=0,
            x1=loadings[i, 0],
            y1=loadings[i, 1]
        )
        fig2.add_annotation(
            x=loadings[i, 0],
            y=loadings[i, 1],
            ax=0, ay=0,
            xanchor="center",
            yanchor="middle",
            text=feature,
        )
    
    return container, fig1, fig2

# Starts app
if __name__ == '__main__':
    app.run_server(debug=True)
    