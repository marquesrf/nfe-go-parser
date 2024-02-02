# NFe (xml format) parser

## This application's goal is to read an NFe (Brazilian digital receipt for online purchased items or services) in xml format and display a sumary with the most important information, such as vendor company name, items(name, quantity, value), taxes, and the total value of the purchase.

**Notice**: this is a WIP (work in progress). Right now, the application expects that a file named NFE.xml is present in its directory to be able to run.

### TODO
- [ ] Check for the presence of any XML file;
- [ ] Receive a XML file to be read as parameter;
- [ ] Read a given directory (as parameter) and read through all xml files;
- [ ] Better print of the samary;
- [ ] Save sumary into a database;
