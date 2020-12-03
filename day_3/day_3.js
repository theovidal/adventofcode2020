const fs = require('fs')

// Read the input file to get the data
fs.readFile('input.txt', 'utf8' , (err, data) => {
  if (err) {
    console.error(`Error while loading the input file: ${err}`)
    return
  }

  rows = data.split('\n')

  // Defining part two tests
  const tests = [
    { right: 1, down: 1 },
    { right: 3, down: 1 }, // Same as part one
    { right: 5, down: 1 },
    { right: 7, down: 1 },
    { right: 1, down: 2 }
  ]

  console.log(`Part 1 - The number of trees encountered are ${testSlope(rows, 3, 1)}`)

  let part2 = tests.reduce((accumulator, test) => accumulator * testSlope(rows, test.right, test.down), 1)
  console.log(`Part 2 - The number is ${part2}`)
})

// testSlope tests descending the slope with specified shifts and returns the number of trees encountered during the trip
function testSlope(rows, rightShift, downShift) {
  let column = 0
  let trees = 0
  for (let row = downShift; row < rows.length; row += downShift) {
    column = (column + rightShift) % 31
    if (rows[row][column] === '#') trees++
  }

  return trees
}
