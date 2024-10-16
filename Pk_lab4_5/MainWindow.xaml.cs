using Microsoft.Win32;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using System.Windows;
using System.Windows.Controls;
using System.Windows.Data;
using System.Windows.Documents;
using System.Windows.Input;
using System.Windows.Media;
using System.Windows.Media.Imaging;
using System.Windows.Navigation;
using System.Windows.Shapes;
using System.IO;
using static System.Net.Mime.MediaTypeNames;
using System.Collections.Specialized;
using System.Diagnostics;
using System.Xml.Linq;


namespace Pk_lab4
{

    public partial class MainWindow : Window
    {
        public MainWindow()
        {
            InitializeComponent();
        }

        private List<string> words = new List<string>();
        private string filePath;
        private string word;
        private int dist;

        private void acceptButton_Click(object sender, RoutedEventArgs e)
        {

            if (filePath!= null)
            {
                Stopwatch time = new Stopwatch();
                time.Start();
                string all_text = File.ReadAllText(filePath);
                char[] separator = {'\n',' ', ',', '.', ':', '\t', '!', '?', '-' };
                
                string[] list = all_text.Split(separator);


                foreach (string s in list)
                {
                    if (words.Contains(s)!= true)
                    {
                        words.Add(s);
                        l_words.Items.Add(new ListBoxItem() { Content = s });
                    }
                }
                time.Stop();
                tm.Text = $"Время загрузки: {time.ElapsedMilliseconds} мс";


            }
        }

        private void escButton_Click(object sender, RoutedEventArgs e)
        {
            this.Close(); // закрытие окна

        }

        private void button1_Click(object sender, EventArgs e)
        {
            OpenFileDialog openFileDlg = new OpenFileDialog();
            openFileDlg.DefaultExt = ".txt";
            openFileDlg.Filter = "Текстовые файлы (*.txt)|*.txt";
            Nullable<bool> result = openFileDlg.ShowDialog();

            if (result == true)
            {
                filePath = openFileDlg.FileName;
                
            }
        }

        

        private void find_w_Click(object sender, RoutedEventArgs e)
        {
            string temp = "Слово не найдено";

            foreach (string s in words)
            {
                if (s == word)
                {
                    temp = $"Слово {word} найдено";
                }   
            }
            enter_w.Text = temp;
          
        }

        private void enter_w_TextChanged(object sender, TextChangedEventArgs e)
        {
            word = enter_w.Text;

        }

        private int Levensh(string w1, string w2)
        {
            int[,] matrix = new int[w1.Length+1, w2.Length+1];
            for( int i = 0; i < w1.Length+1; ++i)
            {
                for ( int j = 0;j < w2.Length + 1; ++j)
                {
                    if (i == 0)
                    {
                        matrix[i, j] = j;
                    }
                    else if (j == 0)
                    {
                        matrix[i, j] = i;
                    }
                    else
                    {
                        matrix[i, j] = 0;
                    }
                }
            }


            for (int i = 1; i < w1.Length+1; ++i)
            {
                for (int j = 1; j < w2.Length+1; ++j)
                {
                    if (w1[i-1] == w2[j-1])
                    {
                        int[] temp_list = { matrix[i - 1, j]+1 , matrix[i, j - 1] +1, matrix[i - 1, j - 1] };
                        matrix[i, j] = temp_list.Min();
                    }
                    else
                    {
                        int[] temp_list = { matrix[i - 1, j] + 1, matrix[i, j - 1] + 1, matrix[i - 1, j - 1] +1};
                        matrix[i, j] = temp_list.Min();
                    }
                }
            }

            return matrix[w1.Length, w2.Length];
        }

        private void dist_Click(object sender, RoutedEventArgs e)
        {
            string temp = "Слово не найдено";

            foreach (string s in words)
            {
                if (Levensh(s,word) <= dist)
                {
                    temp = $"Слово {s}, расстояние Левенштейна {Levensh(s,word)}";
                }
            }
            enter_w.Text = temp;

        }

        private void lev_TextChanged(object sender, TextChangedEventArgs e)
        {
            dist = int.Parse(lev.Text);
        }
    }
}
