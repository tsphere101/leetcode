/**
 * @param {number} x
 * @return {boolean}
 */
 var isPalindrome = function(x) {
    if(x < 0)
        return false;
    myStack = []
    while(x >= 1) {
        myStack.push(x%10);
        x = Math.floor(x/10);
        
    }
    let l = myStack.length;
    for(i = 0 ; i< Math.ceil(l/2); i++)
        {
            if (myStack[i] != myStack[l-1-i])
                return false;
        }
    
    return true;
};