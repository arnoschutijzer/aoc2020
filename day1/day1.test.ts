import { readFile } from '../utilities/parser';
import fixReport from './day1';

describe('day one', () => {
  it('calculates the test case', () => {
    const expenses: number[] = readFile(`${__dirname}/fixtures/day1_test_case.txt`).map(value => parseInt(value));
    expect(fixReport(expenses)).toEqual(514579);
  });
});