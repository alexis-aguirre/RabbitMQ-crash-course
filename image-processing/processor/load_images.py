from typing import List, Iterator
from os import listdir
from os.path import isfile, join
import matplotlib.pyplot as plt
import cv2
import sys
import pytesseract

DATA_PATH = "../data/license-plates"


def get_files_paths(root_path: str = DATA_PATH) -> List[str]:
    return [join(root_path, f) for f in listdir(root_path) if isfile(join(root_path, f))]


def only_jpg_extensions(paths: List[str]) -> Iterator[str]:
    return filter(lambda p: p.endswith('.jpg'), paths)


closed = False


def on_press(event):
    global closed
    sys.stdout.flush()
    if event.key == 'n':
        closed = True
        plt.close()


def main():
    global closed

    for img_path in only_jpg_extensions(get_files_paths()):
        if closed:
            break
        fig, ax = plt.subplots()
        fig.canvas.mpl_connect('key_press_event', on_press)
        img = cv2.imread(img_path)
        resized = cv2.resize(img, (620, 480))
        gray = cv2.cvtColor(resized, cv2.COLOR_BGR2GRAY)
        gray = cv2.bilateralFilter(gray, 13, 15, 15)  # Important to remove noise

        text = pytesseract.image_to_string(gray)
        print("The license plate id is ", text)

        rgb = cv2.cvtColor(gray, cv2.COLOR_GRAY2RGB)
        ax.imshow(rgb)
        plt.show()


main()
