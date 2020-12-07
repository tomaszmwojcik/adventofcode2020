package org.github.tomaszmwojcik;

import java.util.*;
import java.util.function.Predicate;

public class Day4_2 {

    public static boolean isNumberBetween(String v, int len, int min, int max) {
        try {
            int number = Integer.parseInt(v);
            return v.length() == len && number >= min && number <= max;
        } catch(NumberFormatException e) {
            return false;
        }
    }

    private static Set<String> eyeColors = new HashSet<String>(List.of("amb", "blu", "brn", "gry", "grn", "hzl", "oth"));

    public static Map<String, Predicate<String>> validations = Map.ofEntries(
            Map.entry("byr", v -> isNumberBetween(v, 4, 1920, 2002)),
            Map.entry("iyr", v -> isNumberBetween(v, 4, 2010, 2020)),
            Map.entry("eyr", v -> isNumberBetween(v, 4, 2020, 2030)),
            Map.entry("hgt", v -> {
                try {
                    int raw = Integer.parseInt(v.substring(0, v.length() - 2));
                    if (v.endsWith("cm")) {
                        return raw >= 150 && raw <= 193;
                    } else if (v.endsWith("in")) {
                        return raw >= 59 && raw <= 76;
                    }
                    return false;
                } catch(NumberFormatException e) {
                    return false;
                }
            }),
            Map.entry("hcl", v -> v.matches("^#[a-f0-9]{6}$")),
            Map.entry("ecl", v -> eyeColors.contains(v)),
            Map.entry("pid", v -> v.matches("^[0-9]{9}$")),
            Map.entry("cid", v -> true)
    );

    public static void main(String[] args) {
        Set<String> requiredFields = new HashSet<>(Arrays.asList("byr","iyr","eyr","hgt","hcl","ecl","pid"));
        Scanner passportScanner = new Scanner(System.in);
        passportScanner.useDelimiter("\n\n");
        int validPassportsCount = 0;
        while(passportScanner.hasNext()) {
            Scanner recordScanner = new Scanner(passportScanner.next());
            Set<String> readFields = new HashSet<>();
            boolean validationFailed = false;
            while(recordScanner.hasNext()) {
                String[] split = recordScanner.next().split(":");
                readFields.add(split[0]);
                if(!validations.get(split[0]).test(split[1])) {
                    validationFailed = true;
                }
            }
            if(!validationFailed && readFields.containsAll(requiredFields)) {
                validPassportsCount++;
            }
        }
        System.out.println(validPassportsCount);
    }


}
