from django.core.management.base import BaseCommand, CommandError
from bday.models import Key
from bday.settings import BASE_DIR
from random import shuffle
import os


class Command(BaseCommand):
    help = 'Populates database with the randomly scrambled keys from keys.csv.'

    def handle(self, *args, **options):
        keys = []

        # should be done using csv reader but I'm lazy today
        with open(os.path.join(BASE_DIR, 'keys.csv')) as f:
            header = True
            for line in f:
                # skip header
                if header:
                    header = False
                    continue

                # read and parse line
                n, e = list(line.strip().split(','))
                e = int(e)

                # append to keys
                keys.append([n, e])

            shuffle(keys)

            for k in keys:
                n, e = k[0], k[1]

                # create new entry and save it
                k = Key(n=n, e=e)
                k.save()

        self.stdout.write(self.style.SUCCESS('DB populated.'))
