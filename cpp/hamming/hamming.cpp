#include "hamming.h"
#include <iostream>
#include <stdexcept>

using namespace std;

namespace hamming {

int compute(const string& dna1, const string& dna2){
   if(dna1.length() != dna2.length()){
     throw std::domain_error("Strings are different lengths");
    }
    int hammingNum = 0;
    for(unsigned int i = 0; i < dna1.length(); i++){
      if(dna1[i] != dna2[i]){
	hammingNum++;
      }
    }
    return hammingNum;
  }  
}  // namespace hamming
