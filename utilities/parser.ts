import fs from 'fs';
import path from 'path';

export function readFile(filePath: string): string[] {
  const contents = fs.readFileSync(path.resolve(filePath)).toString();
  return contents.split("\n");
}