import re

# Getting the inputs for the exercise
file = open('input.txt', 'r')
passports = file.read().split("\n\n")
passports = list(map(lambda n: n.replace("\n", " ").strip(), passports))

# The height validation function, that is more complex than the others
def hgt(n):
  match = re.compile('[0-9]+(cm|in)').match(n)
  if match is None: return False

  type = match.groups()[0]
  value = n[:-2]
  if type == 'cm': return int(value) in range(150, 193 + 1) # We need to add 1 to the maximum value since it's excluded by the function
  else: return int(value) in range(59, 76 + 1)

# All the validation conditions
validations = {
  'byr': lambda n: int(n) in range(1920, 2002 + 1),
  'iyr': lambda n: int(n) in range(2010, 2020 + 1),
  'eyr': lambda n: int(n) in range(2020, 2030 + 1),
  'hgt': lambda n: hgt(n),
  'hcl': lambda n: re.compile('#[0-9a-f]{6}').match(n) is not None,
  'ecl': lambda n: n in ['amb', 'blu', 'brn', 'gry', 'grn', 'hzl', 'oth'],
  'pid': lambda n: len(str(n)) == 9,
  'cid': lambda _: True
}

part_one_valids = 0
part_two_valids = 0

for passport in passports:
  fields = {}
  
  # Organising passport's fields
  for field in passport.split(' '):
    name, value = field.split(':')
    fields[name] = value

  if len(fields) == 8 or (len(fields) == 7 and not "cid" in fields):
    part_one_valids += 1
    
    valid = True
    # Checking if the passport is valid using fields validations
    for field in fields:
      value = fields[field]
      if not validations[field](value):
        valid = False
        break
    if valid: part_two_valids += 1

print("Part 1 - The number of valid passwords is: {}".format(part_one_valids))
print("Part 2 - The number of valid passwords is: {}".format(part_two_valids))
