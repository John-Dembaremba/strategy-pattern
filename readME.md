# Strategy Pattern

## **Definition**
The **Strategy Pattern** is a behavioral design pattern that enables an algorithm's behavior to be selected at runtime. It defines a family of algorithms, encapsulates each one, and makes them interchangeable, promoting flexibility and reusability of code.

---

## **Key Idea of Strategy Pattern**
The Strategy Pattern allows the behavior of a class to be defined independently of its implementation. By encapsulating algorithms into separate classes and defining a common interface, the pattern enables the class to delegate specific tasks to interchangeable objects, known as strategies.

---

## **Problem it Solves**
The Strategy Pattern addresses the issue of having multiple conditional statements (e.g., `if` or `switch` blocks) to handle varying behaviors or algorithms. These conditionals make the code:

- Difficult to read and maintain.
- Prone to errors when adding new behaviors.
- Violates the **Open/Closed Principle**: the code is not closed for modification when adding new strategies.

By using the Strategy Pattern, you eliminate these conditionals and ensure new behaviors can be added without altering the existing code.

---

## **Components of the Strategy Pattern**
1. **Strategy Interface**:
   - Defines the common interface for all strategies.

2. **Concrete Strategies**:
   - Implement the strategy interface and define specific behaviors or algorithms.

3. **Context**:
   - Maintains a reference to a strategy object and delegates the task to it.

---

## **Use Cases**
- **Payment Gateways**: Selecting between different payment methods like credit card, PayPal, or cryptocurrency.
- **Sorting Algorithms**: Switching between different algorithms such as bubble sort, quicksort, or merge sort based on input size.
- **Authentication**: Supporting multiple authentication mechanisms like OAuth, JWT, or basic authentication.
- **Discount Calculation**: Applying different discount strategies such as percentage-based, fixed value, or no discount.

---

## **Advantages and Disadvantages**
### **Advantages**:
1. **Open/Closed Principle**:
   - Easily add new strategies without modifying existing code.

2. **Single Responsibility Principle**:
   - Separates the logic for various algorithms into their own classes.

3. **Flexibility and Reusability**:
   - Encapsulated strategies are reusable in different contexts.

4. **Runtime Flexibility**:
   - The behavior of the application can change dynamically by selecting different strategies at runtime.

### **Disadvantages**:
1. **Increased Complexity**:
   - Adding multiple strategies may increase the number of classes, making the system more complex.

2. **Overhead**:
   - Introducing abstraction and encapsulation might add a slight runtime overhead.

3. **Code Navigation**:
   - Splitting algorithms into multiple classes can make the code harder to navigate.

---

## **Example Problem**: Discount Calculation
Consider an e-commerce platform where discounts need to be calculated based on different strategies. The strategies could include:

1. **No Discount**: The total amount remains unchanged.
2. **Percentage Discount**: A percentage of the total amount is deducted.
3. **Fixed Discount**: A fixed amount is deducted from the total.

### Why It’s a Good Use Case:
- The discount calculation logic varies depending on the type of promotion.
- Using conditional logic (`if` or `switch` statements) for every discount type leads to hard-to-maintain code.
- New discount types can be added easily by implementing the strategy interface.
- The Strategy Pattern allows dynamic selection of the discount strategy at runtime, enabling personalized promotions.

By using the Strategy Pattern, the discount logic remains clean, maintainable, and extensible.
