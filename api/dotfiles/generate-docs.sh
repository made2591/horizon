java -jar ../docstool/swagger2markup.jar convert -i ../swagger/swagger.json -f ../swagger/index
asciidoctor -D ../docs/ ../swagger/index.adoc