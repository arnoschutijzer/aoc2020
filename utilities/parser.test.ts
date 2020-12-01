import { readFile } from './parser'

describe('parser', () => {
  it('parses some_file.txt', () => {
    expect(readFile(`${__dirname}/fixtures/some_file.txt`)).toMatchSnapshot();
  });
});