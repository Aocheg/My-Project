step 1 
# (1)
# Write pseudocode for a function that checks if a string is a palindrome

    CREATE a function called is_palindrome tha receives strings(str)

        CLEAN the str by removing non alphanumeric charaters and change to lower case

        SET left to 0

        SET right to LENGTH of str MINUS 1

        While left is LESS than right 

            IF character at left in str is not equall to Right then

                Return FALSE 

            Increment left

            Decrement right

        Retun True

    END FUNCTION

# (2)
# Implement your solution in Python.

def is_palindrome (s):
#the above line is defining a function called is_palindrome
    
    s = ''.join(c for c in s if c.isalnum()).lower()
    #the line above clean the strings by removing all non alphanumeric characters and change the cases to lower
    
    left = 0
    #this line uses a pointer to start iterating from the begining which is point zero at the left
    
    right = len(s) - 1
    #this line also uses a pointer to move within the character starting from the right
    
    while left < right:
        #this line is using a loop to continue interating provided the left is less than the right
        
        if s[left] != s[right]:
            #this line shows the real comparison between the left and right if dismatch 
            
            return False
            # if dismach it should return False
            
        left += 1
        # move from the left towards center after a succesful match 
        
        right -= 1
        #move from the right towards centre after a successful match
        
    return True 
    #provided all characters match at the end of the looping return True.


# (3)
# Test examples
print(is_Palindrome("racecar"))
print(is_Palindrome("hello"))
print(is_Palindrome("A man a plan a canal Panama"))


# (4)
# Add comments explaining your reasoning.

    This program uses a two-pointer technique to compare characters from the start and end of the string, ignoring spaces and cases differences, to determine if the string is a Palindrome.

# Step 2

        After running my code through AI i got the following stronger ways to make my code more professional. 
        1. That I added extra memory by creating a strings by using this line "s = ''.join(c for c in s if c.isalnum()).lower()", in a professional way i should have skip the above line and write it in another way that saves memory.
# 3. Reflection
#   What did you learn from solving it before asking AI?

        Before asking AI, I learned how to:
        Break a problem into small logical steps
        Use the two-pointer technique (start and end moving inward)
        Use loops and conditions to control program flow
        Modify a string (remove spaces and change to lowercase)
        Think logically about how characters relate from both ends
        Test my function with different examples 
        Most importantly, I proved to myself that I can solve a problem with my own reasoning.


# How is your understanding different now?
#   After asking AI, my understanding became deeper and more professional:
       
        I now understand time complexity (O(n))
        I learned to think about edge cases (punctuation, symbols, whitespace types)
        I saw that there are multiple ways to solve the same problem
        I learned that my solution is good, but can be improved for real-world use
        I started thinking like a programmer who considers efficiency and robustness, not just “it works”
        I moved from I making it works to I understand why it works and how to make it better. 

# Could you now write similar functions (e.g., reverse a string) without help?
        Yes.
        Because I now understand:
        How to use pointers/indexes
        How to loop through a string
        How to manipulate characters
        How to think step-by-step through a problem
        So writing functions like:
        Reverse a string
        Check if two strings are anagrams
        Count vowels in a string
        Remove duplicates from a string