import requests
from bs4 import BeautifulSoup

# Take user input for URL
url = input("Enter the URL: ")

# Make a request to the website
res = requests.get(url)

# Use BeautifulSoup to parse the HTML
soup = BeautifulSoup(res.text, 'html.parser')

# Initialize empty lists for each file extension
php_links = []
js_links = []
html_links = []
css_links = []
image_links = []
audio_links = []
video_links = []

# Find all the links on the page
for link in soup.find_all('a'):
    # Get the href attribute of the link
    href = link.get('href')
    if href is not None:
        # Check the file extension and add the link to the appropriate list
        if href.endswith('.php'):
            php_links.append(href)
        elif href.endswith('.js'):
            js_links.append(href)
        elif href.endswith('.html'):
            html_links.append(href)
        elif href.endswith('.css'):
            css_links.append(href)
        elif href.endswith('.jpg') or href.endswith('.png') or href.endswith('.gif'):
            image_links.append(href)
        elif href.endswith('.mp3') or href.endswith('.wav'):
            audio_links.append(href)
        elif href.endswith('.mp4') or href.endswith('.avi'):
            video_links.append(href)

# Write the links to a text file
with open('links.txt', 'w') as file:
    file.write("PHP:\n")
    for link in php_links:
        file.write(link + '\n')
    file.write("\nJS:\n")
    for link in js_links:
        file.write(link + '\n')
    file.write("\nHTML:\n")
    for link in html_links:
        file.write(link + '\n')
    file.write("\nCSS:\n")
    for link in css_links:
        file.write(link + '\n')
    file.write("\nImages:\n")
    for link in image_links:
        file.write(link + '\n')
    file.write("\nAudio:\n")
    for link in audio_links:
        file.write(link + '\n')
    file.write("\nVideos:\n")
    for link in video_links:
        file.write(link + '\n')
