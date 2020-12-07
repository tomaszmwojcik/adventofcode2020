# Boarding passes are partioned like FBFBBFFRLR
# First 7 encode seat row (B=1)
# Last 3 - seat(column) (R=1)
# Seat ID: row * 8 + seat
# Find the max seat id in the input

# Part b) find the missing id 
# i.e. the only one that has ID+1 ID-1 present

import sys

def parseRow(encodedRow):
    row = 0
    for c in encodedRow:
        row *= 2
        if c == 'B':
            row += 1
        elif c == 'F':
            pass
        else:
            raise TypeError
    return row

def parseSeat(encodedSeat):
    seat = 0
    for c in encodedSeat:
        seat *= 2
        if c == 'R':
            seat += 1
        elif c == 'L':
            pass
        else:
            raise TypeError
    return seat

if __name__ == "__main__":
    maxSeatId = 0
    seats = []
    for boardingPass in sys.stdin: 
        row = parseRow(boardingPass[0:7])
        seat = parseSeat(boardingPass[7:10]) # sys.stdin left new line at the end
        seatId = row * 8 + seat
        seats.append(seatId)
        maxSeatId = max(seatId, maxSeatId)
    print(maxSeatId)
    sortedSeats = sorted(seats)
    for i in range(1, len(sortedSeats)-1):
        if sortedSeats[i-1]+1 == sortedSeats[i]-1:
            print(sortedSeats[i]-1)
    
