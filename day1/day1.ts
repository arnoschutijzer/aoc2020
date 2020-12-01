import { forEach, map } from 'lodash';
import find from 'lodash/find';
import reduce from 'lodash/reduce';
import sum from 'lodash/sum';

function multiply(...args) {
  return reduce(
    args,
    (prev, curr) => {
      return prev * curr;
    }
  );
}

// TODO: make this generic
export function fixReportPartOne(expenses: number[]) {
  const possibleCombinations = [];

  forEach(expenses, (expense, index) => {
    forEach(expenses.slice(index + 1, expenses.length -1), (secondExpense) => {
      possibleCombinations.push([expense, secondExpense]);
    });
});

  const pair = find([...possibleCombinations], (tuple) => {
    return sum(tuple) === 2020;
  });

  return multiply(...pair);
}

export function fixReportPartTwo(expenses: number[]): number {
  const possibleCombinations = [];

  forEach(expenses, (expense, index) => {
    forEach(expenses.slice(index + 1, expenses.length -1),(secondExpense, secondIndex) => {
      forEach(expenses.slice(secondIndex + 1, expenses.length -1), (thirdExpense) => {
        possibleCombinations.push([expense, secondExpense, thirdExpense]);
      });
    });
  });

  const pair = find([...possibleCombinations], (tuple) => {
    return sum(tuple) === 2020;
  });

  return multiply(...pair);
}