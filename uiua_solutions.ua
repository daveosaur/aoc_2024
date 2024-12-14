# day1
TestOne ← [[3 4] [4 3] [2 5] [1 3] [3 9] [3 3]]

OnePone ← /+⌵/-≡⍆⍉
OnePtwo ← /+×/+⊞=⊣⟜∩⊢.⍉

# day 2
TestTwo ← {[7 6 4 2 1] [1 2 7 8 9] [9 7 6 2 1] [1 3 2 4 5] [8 6 4 4 1] [1 3 6 7 9]}

# helper funcs
# Del        ← ⊂↙⊙(:↘1)⟜↘⊙.
# Del        ← ▽≠⊸⊙(⇡⧻)
Safe       ← /×⊂⊂/×≤3⌵:/×∊°⊂±.↘¯1-↻1.⟜≍⟜◴°□
Del        ← ▽≠⇡⧻,
CheckPerms ← ∊:1≡Safe≡(Del°□₂)≡⊂⇡⊸⧻≡□↯⊸⧻°□

TwoPone ← /+≡Safe
TwoPtwo ← /+≡CheckPerms

# day 3
TestThree ← regex "mul\\((\\d+),(\\d+)\\)"

ThreePone ← /+≡(/×⋕↘1)
