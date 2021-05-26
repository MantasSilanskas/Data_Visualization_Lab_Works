#Lab work Nr.4

import glob
import os
import pandas as pd
from pandas.io.formats import style
import plotly_express as px
import plotly.graph_objects as go
import dash
import dash_core_components as dcc
import dash_html_components as html
from dash.dependencies import Input, Output

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

app = dash.Dash(__name__)

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

cols_to_norm = ['CO2','temperature', 'humidity']
gdf6 = filterDf(df6, "0034C864")
gdf6[cols_to_norm] = gdf6[cols_to_norm].apply(lambda x: (x - x.min()) / (x.max() - x.min()))

gdf6.timestamp = pd.to_datetime(gdf6.timestamp)
gdf6.set_index('timestamp', inplace=True)
grouped = gdf6.groupby(pd.Grouper(freq='1d'))
gdf6 = gdf6[
  ['temperature', 'humidity', 'CO2']
].groupby(pd.Grouper(freq='1d')).mean()

gdf6.reset_index(inplace=True)
gdf6.timestamp = gdf6.timestamp.dt.strftime('%B %d, %Y, %r')

gdf5 = filterDf(df5, "0034C559")
gdf5.timestamp = pd.to_datetime(gdf5.timestamp)
gdf5.set_index('timestamp', inplace=True)
grouped = gdf5.groupby(pd.Grouper(freq='1d'))
gdf5 = gdf5[
  ['temperature', 'humidity', 'CO2']
].groupby(pd.Grouper(freq='1d')).mean()

gdf5.reset_index(inplace=True)
gdf5.timestamp = gdf5.timestamp.dt.strftime('%B %d, %Y, %r')

# App layout
app.layout = html.Div([

html.H1("Data visualization lab work number 4", style={'text-align':'center'}),
html.H2("Chart Nr.1 Line Chart"),

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
html.H2("Chart Nr.2 Box Plot"),
dcc.Graph(id='second_chart', figure={}),

html.Br(),
html.H2("Chart Nr.3 Bar Chart - Max recorded temperature value on each device"),
dcc.Graph(id='third_chart', figure={}),

html.Br(),
html.H2("Chart Nr.4 Scatter plot from October 21th 12:00 pm to October 23th 12:00 pm "),
dcc.Graph(id='fourth_chart', figure={}),

html.Br(),
html.H1("Data visualization lab work number 5", style={'text-align':'center'}),
html.Br(),
html.H2("Chart Nr.1 Scatter plot 3D"),
dcc.Graph(id='fift_chart', figure={}),

html.Br(),
html.H2("Chart Nr.2 Paraller cordinates plot"),
dcc.Graph(id='sixt_chart', figure={}),

html.Br(),
html.H1("Data visualization lab work number 6", style={'text-align':'center'}),
dcc.Graph(id='sevent_chart', figure={}),

])

# Connect dash to plotly

@app.callback(
    [Output(component_id='output_container', component_property='children'),
     Output(component_id='first_chart', component_property='figure'),
     Output(component_id='second_chart', component_property='figure'),
     Output(component_id='third_chart', component_property='figure'),
     Output(component_id='fourth_chart', component_property='figure'),
     Output(component_id="fift_chart", component_property='figure'),
     Output(component_id="sixt_chart", component_property='figure'),
     Output(component_id="sevent_chart", component_property='figure')
     ],
    [Input(component_id='dropDown1', component_property='value'),
    ]
)
def update_graph(option_slctd):
    print(option_slctd)
    print(type(option_slctd))

    container = "The type of data choosen. by user was: {}".format(option_slctd)

    dff1 = df1.copy()
    dff1 = dff1[dff1["type"] == option_slctd]

    dff2 = df2.copy()
    dff2 = dff2[dff2["type"] == option_slctd]

    dff3 = df3.copy()
    dff3 = dff3[dff3["type"] == option_slctd]

    dff4 = df4.copy()
    dff4 = dff4[dff4["type"] == option_slctd]

    dff5 = df5.copy()
    dff5 = dff5[dff5["type"] == option_slctd]

    dff6 = df6.copy()
    dff6 = dff6[dff6["type"] == option_slctd]

    dfs = {"0033C822":dff1,"0034C8DF": dff2, "0034C74E" : dff3, "0034C91E" : dff4, "0034C559" : dff5, "0034C864": dff6}
    
    data = [
        ["0033C822",dff1["value"].max()],["0034C8DF",dff2["value"].max()],
        ["0034C74E",dff3["value"].max()],["0034C91E",dff4["value"].max()],
        ["0034C559",dff5["value"].max()],["0034C864",dff6["value"].max()]
    ]

    df = pd.DataFrame(data, columns = ['device','value'])

    features = ["CO2", "humidity", "temperature"]
    # Plotly Express
    fig1 = go.Figure()
    fig2 = px.box(dff1, y="value")
    fig3 = px.bar(df, x="device", y="value")
    fig4 = px.scatter(dff1[6:132], x="timestamp", y="value")
    fig5 = px.scatter_3d(gdf6[:10], x='CO2', y='temperature', z='humidity', color='timestamp')
    fig6 = px.parallel_coordinates(gdf5[:50], color="CO2", labels={"CO2": "CO2", 
    "humidity": "Humidity","temperature": "Temperature",},
                             color_continuous_scale=px.colors.diverging.Tealrose,
                             color_continuous_midpoint=2)
    fig7 = px.scatter_matrix(
    gdf6,
    dimensions=features,
    color="CO2"
)
    fig7.update_traces(diagonal_visible=False)

    for i in dfs:
     fig1.add_trace(go.Scatter(x = dfs[i]["timestamp"],
                                   y = dfs[i]["value"], 
                                   name = i))
    
    return container, fig1, fig2, fig3, fig4, fig5, fig6, fig7

# Starts app
if __name__ == '__main__':
    app.run_server(debug=True)
    