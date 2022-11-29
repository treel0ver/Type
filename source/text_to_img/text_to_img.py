from PIL import Image, ImageDraw, ImageFont
import sys 

width = 512
height = 512

with open("/home/ggg/Type/source/text_to_img/content") as file:
    lines = [line.rstrip() for line in file]

f = open("/home/ggg/Type/source/text_to_img/content", "r")
ignore = f.readline()

notignore = ' '.join(lines)
notignore = notignore[5:]
new = ""

c = 0
for i in notignore:
    if i == "\u200b":
        new += "\n"
    if c < 25:
        new += i
        c += 1
    else:
        new += i + "\n"
        c = 0

message = "Type's text to image: \n\n" + new

font = ImageFont.truetype("/home/ggg/Type/source/text_to_img/unifont-15.0.01.ttf", size=30)

img = Image.new('RGB', (width, height), color='black')

imgDraw = ImageDraw.Draw(img)

imgDraw.text((10, 10), message, font=font, fill=(255, 255, 255))

img.save('/home/ggg/Type/source/text_to_img/result.png')