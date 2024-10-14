# Notes for self 

- The f(*driver) is used to access the needed functions
- The filepaths that are printed at the start are used by the get sample data function, don't need to touch. 

# Main approach
- Find all files in an organisation 
- Find all the ones with the target directory + something after it
- Print that out
## Reasoning:
- Assuming a correct tree structure, every child/grandchild will have the target directory as a parent.
    - Thus, the parent always will appear in the path before. 
# Progress track by commit
## f86913
- Never used Go before, so needed some time to look through the structure of the package
- Looking at basic IO and calling the target functions.

## 3b58faf
- First try of initial approach