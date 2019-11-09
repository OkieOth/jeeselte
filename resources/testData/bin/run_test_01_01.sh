#!/bin/bash

scriptPos=${0%/*}

stylesheet=$scriptPos/../xslt/test_01_01.xslt
file=$scriptPos/../xml/test_01.xml

xsltproc $stylesheet $file
