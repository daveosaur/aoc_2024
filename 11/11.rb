class Solution
  def initialize(i)
    @inp = i
    @memo = {}
  end

  def solve(target)
    @inp.split(" ").map{|num| parseStones(num.to_i, 1, target) }.sum
  end

  def parseStones(inp, depth, target)
    if depth == target
      return parseSingleStone(inp).length
    end
    return (@memo.key? [inp, depth]) ? 
      @memo[[inp, depth]] :
      parseSingleStone(inp).map{|num|
        @memo[[num, depth+1]] = parseStones(num, depth+1, target) }.sum
  end

  def getDigits(inp)
    left = inp
    right = 0
    len = 0
    multby = 1
    while inp > 0
      inp /= 10
      len+= 1
    end
    (len/2).times do
      right += (left % 10) * multby
      multby *= 10
      left /= 10
    end
    return [left, right, len]
  end

  def parseSingleStone(inp)
    if inp == 0
      return [1]
    else
      nums = getDigits(inp)
      return (nums[2] % 2 == 0) ? [nums[0], nums[1]] : [inp * 2024]
    end
  end
end

inp = File.new("input.txt").read
start = Time.now
puts "part 1: #{Solution.new(inp).solve(25)}"
puts "part 2: #{Solution.new(inp).solve(75)}"
puts (Time.now - start) * 1000

# test
# p Solution.new("125 17").solve(25)
