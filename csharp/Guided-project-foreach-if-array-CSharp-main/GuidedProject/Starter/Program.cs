// initialize variables - graded assignments 
int currentAssignments = 5;

string[] studentNames = ["Sophia", "Andrew", "Emma", "Logan"];

int[] sophiaScores = [90, 86, 87, 98, 100];
int[] andrewScores = [92, 89, 81, 96, 90];
int[] emmaScores = [90, 85, 87, 98, 68,];
int[] loganScores = [90, 95, 87, 88, 96,];

Console.WriteLine("Student\t\tGrade\n");
foreach (string name in studentNames)
{
    if (name == "Sophia")
    {
        int sophiaSum = 0;
        decimal sophiaScore;

        foreach (int score in sophiaScores)
        {
            sophiaSum += score;
        }

        sophiaScore = (decimal)sophiaSum / currentAssignments;

        Console.WriteLine("Sophia:\t\t" + sophiaScore + "\t?");
    }
    if (name == "Andrew")
    {
        int andrewSum = 0;
        decimal andrewScore;

        foreach (int score in andrewScores)
        {
            andrewSum += score;
        }

        andrewScore = (decimal)andrewSum / currentAssignments;

        Console.WriteLine("Andrew:\t\t" + andrewScore + "\t?");
    }
    if (name == "Emma")
    {
        int emmaSum = 0;
        decimal emmaScore;

        foreach (int score in emmaScores)
        {
            emmaSum += score;
        }

        emmaScore = (decimal)emmaSum / currentAssignments;

        Console.WriteLine("Emma:\t\t" + emmaScore + "\t?");
    }
    if (name == "Logan")
    {
        int loganSum = 0;
        decimal loganScore;

        foreach (int score in loganScores)
        {
            loganSum += score;
        }

        loganScore = (decimal)loganSum / currentAssignments;

        Console.WriteLine("Logan:\t\t" + loganScore + "\t?");
    }
}

/* Console.WriteLine("Student\t\tGrade\n");
Console.WriteLine("Andrew:\t\t" + andrewScore + "\tB+");
Console.WriteLine("Emma:\t\t" + emmaScore + "\tB");
Console.WriteLine("Logan:\t\t" + loganScore + "\tA-");
 */
Console.WriteLine("Press the Enter key to continue");
Console.ReadLine();
