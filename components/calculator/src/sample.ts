import { calculator } from '../js-output/calculator.js';

async function main(): Promise<void> {
  console.log('Testing TinyGo Calculator Component via jco transpilation (TypeScript):');
  
  console.log(`10.5 + 5.3 = ${calculator.add(10.5, 5.3)}`);
  console.log(`20.7 - 8.2 = ${calculator.subtract(20.7, 8.2)}`);
  console.log(`4.5 * 3.2 = ${calculator.multiply(4.5, 3.2)}`);
  console.log(`15.6 / 3.9 = ${calculator.divide(15.6, 3.9)}`);
  console.log(`10.0 / 0.0 = ${calculator.divide(10.0, 0.0)}`);
}

main().catch((error: Error) => {
  console.error('Error occurred:', error.message);
  process.exit(1);
});