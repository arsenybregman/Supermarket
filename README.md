# Supermarket

from PyQt5.QtWidgets import QApplication, QLabel, QMainWindow
import pandas as pd

class MainWindow(QMainWindow):
def __init__(self):
super().__init__()

    self.setWindowTitle("My app")
class In_Work_Data():
path = ''
data = pd.DataFrame

def open(self):  # позволяет менять открытый файл
    self.data = pd.read_csv(self.path)

def projection(self, column: str):  # вывод столбца по имени
    return self.data[column]

def value_counts(self, column):  # считает количесвто значений в столбце
    return self.data[column].value_counts()

def clear_column_from_nul(self, column):
    return self.data.dropna(subset=column)
```python
def initialization(path):
    first_data = In_Work_Data()
    first_data.path = path
    first_data.open()
    return first_data

if __name__ == "__main__":
    data = initialization(input())  # здесь Арина принимает инфу с gui
    #print(data.projection('Возраст '))  # и здесь вместо возраста тоже
    #print(data.value_counts('Возраст '))
    app = QApplication([])
    window = MainWindow
    lab = QLabel(data.value_counts('Возраст '))
    lab.show()
    app.exec_()
```
