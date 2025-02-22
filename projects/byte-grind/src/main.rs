mod leetcode;

fn main() {
    println!("Hello ByteGrind");

    let ret = leetcode::p000001::Solution::two_sum(vec![2, 7, 11, 15], 9);
    println!("{:?}", ret);
}
