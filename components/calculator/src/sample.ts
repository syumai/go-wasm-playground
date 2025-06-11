import { calculator } from '../js-output/calculator.js';

interface CalculatorOperation {
  operation: string;
  a: number;
  b: number;
  result: number;
}

async function main(): Promise<void> {
  console.log('Testing TinyGo Calculator Component via jco transpilation (TypeScript):');
  
  const operations: CalculatorOperation[] = [];
  
  // Test addition
  const sum = calculator.add(10.5, 5.3);
  operations.push({ operation: 'addition', a: 10.5, b: 5.3, result: sum });
  console.log(`10.5 + 5.3 = ${sum}`);
  
  // Test subtraction
  const difference = calculator.subtract(20.7, 8.2);
  operations.push({ operation: 'subtraction', a: 20.7, b: 8.2, result: difference });
  console.log(`20.7 - 8.2 = ${difference}`);
  
  // Test multiplication
  const product = calculator.multiply(4.5, 3.2);
  operations.push({ operation: 'multiplication', a: 4.5, b: 3.2, result: product });
  console.log(`4.5 * 3.2 = ${product}`);
  
  // Test division
  const quotient = calculator.divide(15.6, 3.9);
  operations.push({ operation: 'division', a: 15.6, b: 3.9, result: quotient });
  console.log(`15.6 / 3.9 = ${quotient}`);
  
  // Test division by zero
  const divByZero = calculator.divide(10.0, 0.0);
  operations.push({ operation: 'division by zero', a: 10.0, b: 0.0, result: divByZero });
  console.log(`10.0 / 0.0 = ${divByZero}`);
  
  // Summary
  console.log('\nOperation Summary:');
  operations.forEach((op, index) => {
    console.log(`${index + 1}. ${op.operation}: ${op.a} op ${op.b} = ${op.result}`);
  });
}

main().catch((error: Error) => {
  console.error('Error occurred:', error.message);
  process.exit(1);
});