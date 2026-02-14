# Modify your palindrome function to:

    Ignore spaces and punctuation.
    Be case-insensitive.
    Return the position where the string stops being a palindrome (if not one).

def is_palindrome(s): 
    left = 0
    right = len(s) - 1
    
    while left < right: 
       
        if not s[left].isalnum():
            left += 1
        elif not s[right].isalnum():
            right -= 1
       
        elif s[left].lower() != s[right].lower():
            return False, left, right  # return positions where it fails
        else:
            left += 1
            right -= 1

    # Only return True if loop finishes
    return True, None, None

# Test functions
print(is_palindrome("A man, a plan, a canal: Panama"))
print(is_palindrome("hello"))                            
print(is_palindrome("race a car"))   

# The above code is a modified version of my first code (is_palindrome) 
    After attempting to wirte the code and testing it, i got a clearer view from AI that i called .lower() twice per comparison which i used for both lett and right. 

# Reflection

        From this exercise, I learned how to approach a problem step by step and implement a solution on my own. I practiced using loops, conditional statements, and string manipulation to solve a real-world problem. I also learned how to handle different cases, such as ignoring spaces, punctuation, and case differences, to make the function more robust. Most of this knowledge was from after my first palindrome code and getting more insight from AI.
        
        After reviewing my code and thinking about it with AI, I realized that, there are ways to consider edge cases and improve robustness. This experience strengthened my problem-solving skills and gave me confidence in writing similar functions, like reversing a string.