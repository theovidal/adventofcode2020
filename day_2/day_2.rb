inputs = File.open('input.txt').read.split("\n")

part_one = 0
part_two = 0

inputs.each do |input|
  limit, letter, password = input.split
  boundary1, boundary2 = limit.split('-').map(&:to_i)
  letter.chop!
  
  # Part 1 - Boundary 1 is the minimum amount of letters in the password, and Boundary 2 the max
  condensed_password = password.delete("^#{letter}")
  part_one += 1 if condensed_password.length.between? boundary1, boundary2

  # Part 2 - The letter must be located in one out of two spots in the password, defined by the boundaries
  # We substract one to the boundaries in the indexes as there's no index 0 in this challenge
  letter1 = password[boundary1 - 1] == letter
  letter2 = password[boundary2 - 1] == letter
  part_two += 1 if letter1 ^ letter2
end

puts "Part 1 - The amount of valid passwords is #{part_one}"
puts "Part 2 - The amount of valid passwords is #{part_two}"
