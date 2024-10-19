# Notes for self 

- The f(*driver) is used to access the needed functions
- The filepaths that are printed at the start are used by the get sample data function, don't need to touch.
- The more I work on this the more I am realising that go standard should be camel case but I want to get this done ASAP (sorry) 
- Program will execute based on first instance of the directory named in the input
    - I don't think I have the capacity to try and fix the database and clear it of duplicates

- IDEAL SOLUTION: Probably could make this into a tree strucutre for more efficient operations
    - Don't really know how pointers and classes work in go lang, also very little time. 
    - Ideally how I would do it 
        - Make a node struct in folder.go that would be a node for each folder
        - The struct would have a pointer to each child
        - When passing in the data, have a function that converts it to a tree stucture. 
            - Have a helper function here to help read and change the strings. 
        - For GetAllChildren, can just recusively go through 
        - For Move folder, change the pointer and recusively change the strings


# Approach explantions
## Get all children
### High level steps
- Find all files in an organisation 
- Find all the ones with the target directory + something after it
- Print that out
- Runtime: O(n), as we only need to iterate through the whole database once.
    - Should check every file as there is no gurantee that the database is sorted
### Reasoning:
- Assuming a correct tree structure, every child/grandchild will have the target directory as a parent.
    - Thus, the parent always will appear in the path before. 
### Edge cases:
- Duplicate folder names in the same org, especially with both being in different levels of the tree (such as concise cable in the sample data)
    - Solution: Continue to print out children of both
    - Reasoning: Still achieves what is being asked on the task
        - In the case that the duplicates are parent/child, neither will be printed when listing the children 
        - Purpose of the function isn't to fix the database having duplicate names
    - EDIT: there will be no duplicates
- Folders that don't exist or are not in the organisation named are treated the the same. 
    - Reasoning: Organisations exist to separate people's permission levels to access certain content. 
### Analysis 
- Not very efficient 
    - Worst case need to go through the entire database twice  
    - O(n) time

- Ways to makke more efficient
    - Preprocess into tree stucutre, then recusively call 
        - Cna reduce down to O(logn time)

## Move folder 
### High level steps 
- Find file to change and its children and parent
    - If they don't exist then return error
- At the same time, find new parent
    - Error if it doesn't exist
- Do checks 
    - Second folder is not a child of the first 
    - Different orgID 
- Make the required file path changes 
    - Remove from original parent 

### Reasonings 

### Edge cases

### Analysis 
- Not too efficient
    - Need to iterate through every single file to find 
    - Made both implementations do this as there was no guranteed prepocessing for the input. 
        - Maybe could have made my own sort on it, but not too familiar with Go + time. 
# Testing
- Used the testify module as it felt more familiar with testing done before 
- All the data is in the tesitng file itself, admittedly could do better 


# Progress track by commit
## f86913
- Never used Go before, so needed some time to look through the structure of the package
- Looking at basic IO and calling the target functions.

## 3b58faf
- First try of initial approach

## d630ddc
