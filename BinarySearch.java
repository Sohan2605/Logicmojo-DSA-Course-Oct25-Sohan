import java.util.*;


class BinarySearch{
     public static void main(String[] args){
            int[] array = {1 ,2 ,3 ,4 ,5 ,6, 7};
            int search = 6;

            boolean found = binarySearchApply(array,search);
             if (found) {
            System.out.println(search + " found in array!");
        } else {
            System.out.println(search + " not found in array!");
        }
     }

     static boolean binarySearchApply(int[] nums , int search){
           int start = 0;
           int end = nums.length-1;

           while(start<=end){
               int mid = (start + end)/ 2;
               if(nums[mid] == search){
                 return true;
               }else if(search > nums[mid]){
                 start = mid+1;
               }else{
                  end = mid -1;
               }
           }

           return false;
     }
}


// time complexity is logn 
// log16 = ??

    //2^(?) = 16

    //???? means log