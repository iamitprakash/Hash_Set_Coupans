

HashSet<int> usedCoupans = new HashSet<int>();

do
{
    Console.WriteLine("Enter the Coupon Number");
    string couponString = Console.ReadLine();
    if (int.TryParse(couponString, out int coupon))
    {
        if (usedCoupans.Contains(coupon))
        {
            Console.ForegroundColor = ConsoleColor.DarkRed;
            Console.WriteLine("This has been alreday used :-(");
            Console.ForegroundColor = ConsoleColor.Gray;

        }
        else
        {
            usedCoupans.Add(coupon);
            Console.ForegroundColor = ConsoleColor.Green;
            Console.WriteLine("Thank You! :-)");
            Console.ForegroundColor = ConsoleColor.Gray;
        }
    }
    else
    {
        break;
    }

}
while (true);
