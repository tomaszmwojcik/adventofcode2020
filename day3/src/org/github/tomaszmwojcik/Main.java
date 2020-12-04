package org.github.tomaszmwojcik;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;
import java.util.Scanner;

/*
Count trees denoted by '#' when moving [dx fields right, dy down] starting from top-left
 */
public class Main {
    private static int checkSlope(List<String> input, int dx, int dy) {
        int x = 0;
        int treeCount = 0;
        int xMax = input.get(0).length();
        for (int y = 0; y < input.size(); y += dy) {
            if (input.get(y).charAt(x % xMax) == '#') {
                treeCount++;
            }
            x += dx;
        }
        System.out.println(treeCount);
        return treeCount;
    }

    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        List<String> input = new ArrayList<>();
        while (scanner.hasNextLine()) {
            input.add(scanner.nextLine());
        }
        System.out.printf("When moving [3,1] you'll encounter %d trees %n", checkSlope(input, 3, 1));
        int[][] params = {{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}};
        long result = Arrays.stream(params).map(paramPair -> (long) checkSlope(input, paramPair[0], paramPair[1])).reduce(1L, (a, b) -> a * b);
        System.out.println("Second star password is " + result);
    }
}
