
Console.Write("Введите коэффицент A\n");
double a = Convert.ToDouble(Console.ReadLine());
Console.Write("Введите коэффицент B\n");
double b = Convert.ToDouble(Console.ReadLine());
Console.WriteLine("Введите коэффицент C\n");
double c = Convert.ToDouble(Console.ReadLine());

while (a == 0 && b == 0 && c != 0)
{
    Console.Write("Введите новые коэффиценты\n");
    Console.Write("Введите коэффицент A\n");
    a = Convert.ToDouble(Console.ReadLine());
    Console.Write("Введите коэффицент B\n");
    b = Convert.ToDouble(Console.ReadLine());
    Console.WriteLine("Введите коэффицент C\n");
    c = Convert.ToDouble(Console.ReadLine());
}
Uravn n1 = new Uravn(a, b, c);
n1.Print();
Uravn n2 = new Uravn(1, 2, 1);
n2.Print();


class Uravn
{
    public double a;
    public double b;
    public double c;

    public Uravn() { a = 0; b = 0; c = 0; }
    public Uravn(double p1, double p2, double p3) { a = p1; b = p2; c = p3; }

    public double Find_disc()
    {
        return Math.Pow(b, 2) - 4 * a * c;
    }

    public void Print()
    {
        Console.ForegroundColor = ConsoleColor.Green;

        if (a == 0 && b == 0 && c == 0)
        {
            Console.WriteLine("Бесконечное множество корней\n");
        }
        else if (a == 0 && b != 0 && c != 0)
        {
            double x1 = -c / b;
            Console.WriteLine($"Корень: {x1}\n");
        }
        else
        {
            double d = Find_disc();
            if (d == 0)
            {
                double x1 = -b / (2 * a);
                Console.WriteLine($"Корень: {x1}\n");
            }
            else if (d > 0)
            {
                double x1 = (-b + Math.Sqrt(d)) / (2 * a);
                double x2 = (-b - Math.Sqrt(d)) / (2 * a);
                Console.WriteLine($"Корни: {x1},{x2}\n");
            }
            else
            {
                Console.ForegroundColor = ConsoleColor.Red;
                Console.WriteLine("Нет корней");
            }
        }
        Console.ResetColor();
    }
}



