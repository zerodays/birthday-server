from django.db import models


class Key(models.Model):
    n = models.CharField(max_length=1000, null=False)
    e = models.IntegerField(null=False)
