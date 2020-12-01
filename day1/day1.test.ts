import { readFile } from '../utilities/parser';
import { fixReportPartOne, fixReportPartTwo } from './day1';

describe('day one', () => {
  it('calculates the test case', () => {
    const expenses: number[] = readFile(`${__dirname}/fixtures/day1_test_case.txt`).map(value => parseInt(value));
    expect(fixReportPartOne(expenses)).toEqual(514579);
  });

  it('calculates the multiplication of the tuple with size 3', () => {
    const expenses: number[] = readFile(`${__dirname}/fixtures/day1_test_case.txt`).map(value => parseInt(value));
    expect(fixReportPartTwo(expenses)).toEqual(241861950);
  });

  it('calculates the multiplication of the tuple with size 2', () => {
    const expenses: number[] = readFile(`${__dirname}/fixtures/day1_input.txt`).map(value => parseInt(value));
    expect(fixReportPartOne(expenses)).toEqual(805731);
  });

  it('calculates the multiplication of the tuple with size 3', () => {
    const expenses: number[] = readFile(`${__dirname}/fixtures/day1_input.txt`).map(value => parseInt(value));
    expect(fixReportPartTwo(expenses)).toEqual(192684960);
  });
});