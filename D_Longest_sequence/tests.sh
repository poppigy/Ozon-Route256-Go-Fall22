for t in $(ls tests | grep -v ".a"); do cat "tests/${t}" | ./D_Longest_sequence > "${t}.a" && diff -y "tests/${t}.a" "${t}.a" > "${t}_diff"; done
