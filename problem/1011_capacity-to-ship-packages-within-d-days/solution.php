<?php
// [1011. 在 D 天内送达包裹的能力](https://leetcode-cn.com/problems/capacity-to-ship-packages-within-d-days/)

class Solution {

    /**
     * @param array[] $weights
     * @param Integer $D
     * @return Integer
     */
    function shipWithinDays($weights, $D) {
        $L_left= max($weights);
        $L_right = array_sum($weights);

        while($L_left != $L_right){
 
            $L_mid = $L_left + (int)(($L_right - $L_left) / 2);

            for($i=$w_sum=0, $count_D = 1; $i < count($weights); $i++){
                $w_sum += $weights[$i];
                if ($w_sum > $L_mid){
                    $w_sum = $weights[$i];
                    $count_D++;
                }
            }

            $count_D <= $D ? $L_right = $L_mid : $L_left = $L_mid + 1;
            
        }

        // $L_mid = intval(round($L_right + $L_left) / 2);
        // echo "===================\n";
        // echo "L_左：\t".$L_left."\n";
        // echo "L_中：\t".$L_mid."\n";
        // echo "L_右：\t".$L_right."\n\n";

        return $L_left;

    }
}

$weights = [
    3 => [3,2,2,4,1,4],
    4 => [1,2,3,1,1],
    5 => [1,2,3,4,5,6,7,8,9,10],    
];
$test = new Solution();

echo $test -> shipWithinDays($weights[3], 3)."\n";
echo $test -> shipWithinDays($weights[4], 4)."\n";
echo $test -> shipWithinDays($weights[5], 5)."\n";