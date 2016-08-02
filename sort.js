process.stdin.resume();
process.stdin.setEncoding('ascii');

var input_stdin = "";
var input_stdin_array = "";
var input_currentline = 0;

process.stdin.on('data', function (data) {
    input_stdin += data;
});

process.stdin.on('end', function () {
    input_stdin_array = input_stdin.split("\n");
    main();    
});

function readLine() {
    return input_stdin_array[input_currentline++];
}

function bubble(a, n) {
  var output = [];
  var totalSwaps = 0;
  for (var i = 0; i < n; i++) {
    var numberOfSwaps = 0;
      
      for (var j = 0; j < n - 1; j++) {
          if (a[j] > a[j + 1]) {
              // swap
              var temp = a[j];
              a[j] = a[j+1];
              a[j+1] = temp;
              numberOfSwaps++;
              totalSwaps++;
          }
      }
      
      if (numberOfSwaps == 0) {
          break;
      }
  }
  output[0] = totalSwaps;
  output[1] = a[0];
  output[2] = a[n-1];
  return output;
}

function main() {
    var n = parseInt(readLine());
    a = readLine().split(' ');
    a = a.map(Number);
    var output = bubble(a, n);
    console.log("Array is sorted in " + output[0] +" swaps.");
    console.log("First Element: " + output[1]);
    console.log("Last Element: "+ output[2]);


}
