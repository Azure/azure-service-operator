# Package Documentation Instructions

Every Go package in ASO should have a `README.md` file that provides essential information about the package for maintainers and contributing developers. 

The structure of the `README.md` file should follow the structure of this document, with each section containing the indicated content.

If you have any suggestions on additional content, please discuss by raising questions, but otherwise don't add any additional content not covered by this template.

## Overview

First, a brief - two or three sentences - description of the package, its purpose, and its main features.

**Key concepts**: After the summary, a small number of paragraphs, each describing a key concept needed for understanding of the package. This should be a high-level overview of the concepts, not a detailed explanation. Each concept paragraph should follow the format of this paragraph, starting with a bolden phrase that summarizes the concept, followed by a more detailed explanation. Don't use a single paragraph called "Key Concepts", this is an example.

## Testing

A brief description of how to test the package, including any specific commands or scripts that should be run. 

Be sure to include any additional information you think relevant to developers needing to test the package, including, but not limited to:

* Which test libraries are used for assertions, mocking, golden tests, recordings, etc etc.
* How to set up the test environment (e.g. required environment variables)

## Related packages

A reference list of the most used packages within this repo that are either referenced by this package, or which refer to this package. This should include a brief description of each package and its relationship to the current package.



