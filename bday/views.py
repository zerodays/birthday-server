from django.shortcuts import render, HttpResponse
from django.db import connection
import json


def run_sql_query(request):
    q = request.GET['q']
    cursor = connection.cursor()
    cursor.execute(q)
    result = []
    for i in cursor.fetchall():
        result.append(i)
    return HttpResponse(json.dumps(result))


def index(request):
    return render(request, 'bday/index.html')
