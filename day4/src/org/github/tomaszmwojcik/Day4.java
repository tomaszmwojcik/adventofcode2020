package org.github.tomaszmwojcik;

import java.util.Arrays;
import java.util.HashSet;
import java.util.Scanner;
import java.util.Set;

public class Day4 {

    public static void main(String[] args) {
        Set<String> requiredFields = new HashSet<>(Arrays.asList("byr","iyr","eyr","hgt","hcl","ecl","pid"));
        Scanner passportScanner = new Scanner(System.in);
        passportScanner.useDelimiter("\n\n");
        int validPassportsCount = 0;
        while(passportScanner.hasNext()) {
            Scanner recordScanner = new Scanner(passportScanner.next());
            Set<String> readFields = new HashSet<>();
            while(recordScanner.hasNext()) {
                readFields.add(recordScanner.next().split(":")[0]);
            }
            if(readFields.containsAll(requiredFields)) {
                validPassportsCount++;
            }
        }
        System.out.println(validPassportsCount);
    }


}
