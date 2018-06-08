"""
    Test plotly
"""

import plotly.graph_objs as go
import plotly.offline as offline

def interpret_equation(string):
    """
    Get a and b from a string that has the form y = a x + b, 
    where each character is separated by a space.

    TODO: Build a more generic function that can interpret different operator and also different
    type of functions (not necessarily linear)
    """
    nums = [int(s) for s in string.split() if s.isdigit()]
    return nums[0], nums[1]

def linear_function(a, b, x):
    """
    Return y = ax + b.
    """
    return (a * x) + b

def build_up_value(n, a, b):
    """
    Build a list of values for x and y to make a graph.

    TODO: Only works for linear_function written above. Can be extended further.
    """
    x = []
    y = []

    for i in range(0, n):
        x.append(i)

    for i in x:
        r = linear_function(a=a, b=b, x=i)
        y.append(r)

    return x, y

def plot(list_x, list_y, title):
    """
    Plot a graph base on multiple values of x and y.

    TODO: 
        - Check if x and y has the same length.
        - Make a more complex graph if needed.
    """
    trace1 = go.Scatter(x=list_x, y=list_y, marker={'color': 'red', 'symbol': 104, 'size': "10"}, 
                        mode="markers+lines")
                                                   
    data=go.Data([trace1])
    layout=go.Layout(title=title, xaxis={'title':'x'}, yaxis={'title':'y'})
    figure=go.Figure(data=data,layout=layout)
    offline.plot(figure, filename='{0}simple.html'.format('./static/'))

if __name__ == '__main__':
    # Sample valid equation
    sample_equation = 'y = 2 x + 1'

    # Get a and b
    a, b = interpret_equation(sample_equation)

    # Graph y = ax + b using 10 values
    x, y = build_up_value(n=10, a=a, b=b)
    plot(x, y, 'Linear Function')
