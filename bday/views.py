from django.shortcuts import render


def run_sql_query(request):
    pass


def index(request):
    return render(request, 'bday/index.html')
