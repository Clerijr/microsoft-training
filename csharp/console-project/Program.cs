string orderStream = "B123,C234,A345,C15,B177,G3003,C235,B179";
string[] orderArr = orderStream.Split(',');
Array.Sort(orderArr);

foreach (string order in orderArr)
{
    if (order.Length != 4)
    {
        Console.WriteLine($"{order}: {order.Length} - Error");
    } else {
        Console.WriteLine($"{order}: {order.Length}");
    }
}