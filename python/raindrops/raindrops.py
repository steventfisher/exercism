def convert(number):
    if(number % 3 == 0):
        if(number % 5 == 0):
            if(number % 7 == 0):
                return 'PlingPlangPlong'
            return 'PlingPlang'
        elif(number % 7 == 0):
            return 'PlingPlong'
        return 'Pling'
    elif(number %5 == 0):
        if(number % 7 == 0):
            return 'PlangPlong'
        return 'Plang'
    elif(number % 7 == 0):
        return 'Plong'
    else:
        return str(number)
