import csv
import pandas as pd
import seaborn as sns
import matplotlib as plt


def isiterable(obj):
    try:
        iter(obj)
        return True
    except TypeError:  # не является итерируемым
        return False


class In_Work_Data():
    def __init__(self, dir: str):
        self.path = dir
        self.data = pd.read_csv(dir)

    def projection(self, column: str):  # вывод столбца по имени
        return self.data[column]

    def value_counts(self, column: str):  # считает количесвто значений в столбце
        return self.data[column].value_counts()

    def clear_column_from_nul(self, column):
        return self.data.dropna(subset=column)

    def change_nul_for_x(self, column: str, x):
        _ = self.data.fillna()
        return self.data

    def first_five(self):
        for i in self.data.columns.tolist():
            self.first_five_by_column(i)

    def popular(self):
        for i in self.data.columns.tolist():
            print(i)
            print(pd.Series(self.data[i].value_counts()))

    def first_five_by_column(self, i):
        line = {}
        # print(pd.Series(self.data.value_counts(i))['Хлеб'])
        for j in self.data[i].unique():
            #  if j in pd.Series(self.value_counts(i)).index:
            if type(j) == float:
                j = 'Не покупаю'
            mnozitel = pd.Series(self.data.value_counts(i))[j]
            for k in j.split(';'):
                if k not in list(line.keys()):
                    line[k] = 1
                elif k in list(line.keys()):
                    line[k] = line[k] + 1 * mnozitel
        print(i)
        print(line)


# class First_quiz()
#    def __init__(self):
#        self.point = {'За продуктами': 0, 'За перекусом'}

if __name__ == '__main__':
    new_data = In_Work_Data(input())  # здесь Арина принимает инфу с gui
    #  print(new_data.value_counts('Какую хлебобулочную продукцию вы покупаете?'))  # и здесь вместо возраста тоже
    # print(new_data.value_counts('Цель посещения продуктового магазина?'))
    # new_data.popular()
    new_data.first_five_by_column("Какую хлебобулочную продукцию вы покупаете?")
    new_data.first_five_by_column("Какую мясную продукцию вы покупаете?")
    # print(new_data.first_five())
