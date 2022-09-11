for t in $(ls tests | grep -v ".a"); do cat "tests/${t}" | ./F_Opposite_elements > "${t}.a" && diff -y "tests/${t}.a" "${t}.a" > "${t}_diff"; done
