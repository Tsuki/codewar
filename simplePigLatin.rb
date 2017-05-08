# https://www.codewars.com/kata/520b9d2ad5c005041100000f/train/ruby

def pig_it text
  text.split().map{|v| v =~ /[[:punct:]]/  ? v : v.split('').rotate(1).join()+"ay"}.join(' ')
end
